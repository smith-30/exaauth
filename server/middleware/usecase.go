package middleware

// var (
// 	UsecaseWithRDB = &contextKey{"UsecaseWithRDB"}
// )

// func UsecaseWithRDB() func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		fn := func(w http.ResponseWriter, r *http.Request) {
// 			tx := orm.DBConn().Begin()
// 			ctx := r.Context()
// 			ctx = context.WithValue(ctx, TxKey, tx)
// 			ctx = context.WithValue(ctx, DoCommit, false)

// 			next.ServeHTTP(w, r)

// 			if ctx.Value(DoCommit).(bool) {
// 				tx.Commit()
// 			} else {
// 				tx.Rollback()
// 			}
// 		}
// 		return http.HandlerFunc(fn)
// 	}
// }
