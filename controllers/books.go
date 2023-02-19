package controllers

// func (s *Server) Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		uname, pwd, ok := r.BasicAuth()
// 		if !ok {
// 			w.Write([]byte("Username atau Password tidak boleh kosong"))
// 			return
// 		}

// 		if uname == "admin" && pwd == "password" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}else if uname == "editor" && pwd == "secret" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		w.Write([]byte("Username atau Password tidak sesuai"))
// 		return
// 	})
// }
