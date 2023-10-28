package transport

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	connect_go "github.com/bufbuild/connect-go"
	pkgerrors "github.com/pkg/errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1alpha1 "ddd-sample/interface/transport/service/v1alpha1"
	"ddd-sample/interface/transport/service/v1alpha1/svcv1alpha1connect"
	typesv1alpha1 "ddd-sample/interface/transport/types/v1alpha1"
	"ddd-sample/usecase"
)

type IServer interface {
	Serve(ctx context.Context, l net.Listener) error
}

const (
	GracefulShutdownSeconds = 3 * time.Second
)

type Server struct {
	ticketUseCase usecase.ITicketUsecase
}

var (
	_ svcv1alpha1connect.DDDSampleServiceHandler = (*Server)(nil)
	_ IServer                                    = (*Server)(nil)
)

func NewServer(ticketUseCase usecase.ITicketUsecase) IServer {
	return &Server{
		ticketUseCase: ticketUseCase,
	}
}

func (s Server) Serve(ctx context.Context, l net.Listener) error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Handler: h2c.NewHandler(mux,
			&http2.Server{}),
		ReadHeaderTimeout: time.Minute,
	}
	// construct handler and path
	// can we wire it?
	path, svcHandler := svcv1alpha1connect.NewDDDSampleServiceHandler(s)
	mux.Handle(path, svcHandler)

	errCh := make(chan error, 2)
	go func() { errCh <- srv.Serve(l) }()
	log.Printf("IServer is listening on %q", l.Addr().String())
	select {
	case <-ctx.Done():
		log.Println("Gracefully stopping Server...")
		time.Sleep(GracefulShutdownSeconds)
		return srv.Shutdown(context.Background())
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (s Server) CreateTicket(ctx context.Context, c *connect_go.Request[v1alpha1.CreateTicketRequest]) (*connect_go.Response[v1alpha1.CreateTicketResponse], error) {
	id, err := s.ticketUseCase.Create(c.Msg.GetContent(), c.Msg.GetStart().AsTime(), c.Msg.GetEnd().AsTime())
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, pkgerrors.Wrap(err, "create ticket"))
	}
	return connect_go.NewResponse(&v1alpha1.CreateTicketResponse{
		Id: id,
	}), nil
}

func (s Server) GetTicket(ctx context.Context, c *connect_go.Request[v1alpha1.GetTicketRequest]) (*connect_go.Response[v1alpha1.GetTicketResponse], error) {
	useCaseTicket, err := s.ticketUseCase.Get(c.Msg.GetId())
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, pkgerrors.Wrap(err, "get ticket"))
	}
	return connect_go.NewResponse(&v1alpha1.GetTicketResponse{
		Ticket: toTransportTicket(&useCaseTicket),
	}), nil
}

func toTransportTicket(in *usecase.Ticket) (out *typesv1alpha1.Ticket) {
	return &typesv1alpha1.Ticket{
		// getter for usecase ticket?
		Content: in.Content,
		Id:      in.Id,
		Start:   timestamppb.New(in.ActiveWindow.Start),
		End:     timestamppb.New(in.ActiveWindow.End),
	}
}
