package main
import (
"encoding/json"
"fmt"
"net/http"
"time"

"github.com/hashicorp/vault/api"
"github.com/hashicorp/vault/builtin/credential/aws"
)

const (
vaultAddr = "http://YOUR_VAULT_ADDR:8200"

staticToken = "YOUR_STATIC_TOKEN"
)

var httpClient = &http.Client{
Timeout: 10 * time.Second,
}

func main() {
// token := staticToken
//token, err := userpassLogin()
//if err != nil {
//	panic(err)
//}
token, err := awsLogin()
if err != nil {
panic(err)
}

client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
if err != nil {
panic(err)
}

client.SetToken(token)
data, err := client.Logical().Read("secret/data/my-secret")
if err != nil {
panic(err)
}

b, _ := json.Marshal(data.Data)
fmt.Println(string(b))
}

const (
username = "jun06t"
password = "foo"
)

func userpassLogin() (string, error) {
// create a vault client
client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
if err != nil {
return "", err
}

// to pass the password
options := map[string]interface{}{
"password": password,
}
path := fmt.Sprintf("auth/userpass/login/%s", username)

// PUT call to get a token
secret, err := client.Logical().Write(path, options)
if err != nil {
return "", err
}

token := secret.Auth.ClientToken
return token, nil
}

const (
accessKey    = ""
secretKey    = ""
sessionToken = ""
headerValue  = ""
)

func awsLogin() (string, error) {
// get aws credential
data, err := awsauth.GenerateLoginData(accessKey, secretKey, sessionToken, headerValue)
if err != nil {
return "", err
}

// create a vault client
client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
if err != nil {
return "", err
}

// PUT call to get a token
secret, err := client.Logical().Write("auth/aws/login", data)
if err != nil {
return "", err
}

token := secret.Auth.ClientToken
return token, nil
}