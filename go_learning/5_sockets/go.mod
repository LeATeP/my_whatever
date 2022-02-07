module leatep/sockets

go 1.17

replace leatep/handler => ./clientHandler

require (
	leatep/handler v0.0.0-00010101000000-000000000000
	leatep/pdb v0.0.0-00010101000000-000000000000
)

require github.com/lib/pq v1.10.4 // indirect

replace leatep/pdb => ../8_db_module
