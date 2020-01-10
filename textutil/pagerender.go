package textutil

import "fmt"

//import "net/http"

const pageHead string = `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">

  <title>URL Shortener</title>
  <meta name="description" content="A Simple URL Shortener Written In Go">
  <meta name="author" content="V. Timchenko">
  <link rel="stylesheet" 
  href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" 
  integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" 
  crossorigin="anonymous">
</head>

<body><div class="container">`

const pageForm string = `<form>
<div class="form-group">
  <label for="url">URL</label>
  <input type="url" class="form-control" id="url" placeholder="Enter URL">
</div>
<button type="button" class="btn btn-primary" onclick="sendPost();">Shorten!</button>
</form>`

const jsCode = `<script>
function sendPost(){ const url = 'http://%s/';
const Http = new XMLHttpRequest();
var data = {}
data.URL = document.getElementById("url").value
Http.open("POST", url);
Http.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
Http.onreadystatechange = function(){
  if (this.readyState == 4 && this.status == 200) {
    var result = JSON.parse(this.responseText)
    alert = document.getElementById("alert")
    if ("error" in result) {
      alert.className = "alert alert-danger"
      alert.innerHTML = result.error;
    } else {
      alert.className = "alert alert-success"
      alert.innerHTML = result.success;
    }
    alert.style.display = "block";
    return;
  }
};
Http.send(JSON.stringify(data));
}</script>`

const pageFooter string = `</div></body>
</html>
`

func generateAlert(msg, role string) string {
	var display string
	if msg == "" {
		display = "none"
	} else {
		display = "block"
	}
	return fmt.Sprintf("<div id=\"alert\" class=\"alert alert-%s\" role=\"alert\" style=\"display:%s\">%s</div>", role, display, msg)
}

func generateJSCode(host string) string {
	return fmt.Sprintf(jsCode, host)
}

// RenderResponse generates the rendered HTML code to be returned to browser in
// response to the given request, with the addition of the error or success messages
func RenderResponse(host, errMsg string) string {
	return pageHead +
		generateAlert(errMsg, "danger") +
		pageForm +
		generateJSCode(host) +
		pageFooter
}
