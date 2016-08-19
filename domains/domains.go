package domains

type Email struct {
	To      string
	Subject string
	Body    string
}

type ServerConfig struct {
	Login struct {
		Glogin string
		
	}
	Pass struct {
		Gpass string
		
	}

	// Dbmgo struct {
	// 	Addrs     []string
	// 	Database  string
	// 	Username  string
	// 	Password  string
	// 	Mechanism string
	// }

}
