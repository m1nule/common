package response

import (
	"net/http"

	"github.com/m1nule/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	w.Header().Set(trace.TraceIdKey, trace.TraceIDFromContext(r.Context()))
	retCode := http.StatusOK
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, retCode, r)
	} else {
		retCode = http.StatusBadRequest
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器错误"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				retCode = getHttpCode(gstatus.Code())
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[API-ERR] : %+v ", err)

		httpx.WriteJson(w, retCode, Error(errcode, errmsg))
	}
}

// 授权的http方法
// func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
// 	w.Header().Set(trace.TraceIdKey, trace.TraceIDFromContext(r.Context()))
// 	if err == nil {
// 		// 成功返回
// 		r := Success(resp)
// 		httpx.WriteJson(w, http.StatusOK, r)
// 	} else {
// 		// 错误返回
// 		errcode := xerr.SERVER_COMMON_ERROR
// 		errmsg := "服务器错误"
//
// 		causeErr := errors.Cause(err)                // err类型
// 		if e, ok := causeErr.(*xerr.CodeError); ok { // 自定义错误类型
// 			// 自定义CodeError
// 			errcode = e.GetErrCode()
// 			errmsg = e.GetErrMsg()
// 		} else {
// 			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
// 				grpcCode := uint32(gstatus.Code())
// 				if xerr.IsCodeErr(grpcCode) {
// 					errcode = grpcCode
// 					errmsg = gstatus.Message()
// 				}
// 			}
// 		}
//
// 		logx.WithContext(r.Context()).Errorf("[GATEWAY-ERR] : %+v ", err)
//
// 		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
// 	}
// }

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error, mode string) {
	w.Header().Set(trace.TraceIdKey, trace.TraceIDFromContext(r.Context()))
	if mode != service.ProMode {
		logx.WithContext(r.Context()).Errorf("[API-ERR] [参数错误]: %+v ", err.Error())
	}
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR)))
}

func getHttpCode(gCode codes.Code) int {
	switch gCode {
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.NotFound:
		return http.StatusNotFound
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	default:
		return http.StatusBadRequest
	}
}
