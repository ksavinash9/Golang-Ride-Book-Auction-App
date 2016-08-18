package config

//Set the configuration values on this file

// Get gets you the value for the key
func Get(t string) string {
    switch t {
    case "WEBSERVER_PORT":
        return "9000"
    case "DATABASE_NAME":
        return "database"
    case "DATABASE_USER":
        return "swarn"
    case "SSL_MODE":
        return "disable"
    default:
        return ""
    }
}
