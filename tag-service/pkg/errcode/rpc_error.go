package errcode

import (
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	pbErr := &pb.Error{
		Code: int32(err.Code()),
		Message: err.Msg(),
	}
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(pbErr)
	return s.Err()
}

// 将错误码转换为对应的gRPC内部定义的错误码
func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}

// 获取错误类型
type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

// 除了希望把原始业务错误码放进Details中, 还想放进其他信息的时候可以用该方法
func ToRPCStatus(code int, msg string) *Status {
	pbErr := &pb.Error{
		Code: int32(code),
		Message: msg,
	}
	s, _  := status.New(ToRPCCode(code), msg).WithDetails(pbErr)
	return &Status{
		s,
	}
}
