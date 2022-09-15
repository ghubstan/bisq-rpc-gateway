package proxy

import (
	"encoding/json"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "golang.org/x/bisq-grpc-gateway/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type errorBody struct {
	Err string `json:"error,omitempty"`
}

// Maps gRPC status to HTTP status.
func BisqHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Err: status.Convert(err).Message(),
	})
	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

func RunProxy(proxyAddress string, serverAddress string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	runtime.HTTPError = BisqHTTPError
	mux := runtime.NewServeMux(opts...)
	prettier := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Checks Values as map[string][]string also catches ?pretty and ?pretty=
			// r.URL.Query().Get("pretty") would not.
			if _, ok := r.URL.Query()["pretty"]; ok {
				r.Header.Set("Accept", "application/json+pretty")
			}
			h.ServeHTTP(w, r)
		})
	}

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := flag.String("one_size_fits_all_endpoint", serverAddress, "todo")
	err := pb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, *endpoint, dialOpts)
	// TODO Fix:  error is nil even when gRPC server is not running.
	if err != nil {
		return err
	}

	log.Printf("RPC proxy registered at %s will forward requests to %s\n", proxyAddress, serverAddress)

	// Set HTTP server to listen for incoming requests
	log.Fatal(http.ListenAndServe(proxyAddress, prettier(mux)))
	return nil
}
