package main //影像发布处理

import (
	"fmt"
	"ld_yjt_img_http/Ld_yjt_dir_http"
	"ld_yjt_img_http/Ld_yjt_file_http"
	"ld_yjt_img_http/Ld_yjt_log"
	"net/http"
)

const (
	DEFINE_FILE = "ld_yjt_file"
	DEFINE_DIR  = "ld_yjt_dir"
)

func main() {
	hostAndPort := ":12888"

	Ld_yjt_log.WriteLog(fmt.Sprintf("开启文件发布服务成功,参数为[%s]", hostAndPort), true)
	http.HandleFunc("/"+DEFINE_FILE+"/", Ld_yjt_file_http.HttpFileSvr)
	http.HandleFunc("/"+DEFINE_DIR+"/", Ld_yjt_dir_http.HttpDirSvr)

	http.ListenAndServe(hostAndPort, nil)
}
