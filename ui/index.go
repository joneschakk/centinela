package ui

const IndexHTML = `

<div class="form">
  <input type="text" class="text-box" id="in-username" placeholder="Username" />
  <input type="password" class="text-box" id="in-password" placeholder="Password" />
  <input type="submit" class="btn-submit" value="Login" onclick="submit()" />
  <label id="message-label" class="messaage" style="hidden:true"></label>
  <!-- <a href="">No credentials?</a> -->
</div>

<script>
  var target = "{{AuthTarget}}"
  var targetUrl = "{{TargetUrl}}"

  function showMessage(msg) {
    var lbl = document.getElementById('message-label')
    lbl.innerText = msg
    lbl.hidden = false
  }

  function submit() {
    var username = document.getElementById("in-username").value
    var password = document.getElementById("in-password").value
    // const data = new FormData();
    // data['username'] = username
    // data['password'] = password
    // console.log(data);
    domain = new URL(targetUrl).host
    fetch("/authenticate", {
      method: 'post',
      body: "target=" + target + "&domain=" + domain + "&username=" + username + "&password=" + password,
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      }
    }).then(
      resp => {
        console.log("resp", resp)
        if (resp.ok) {
          alert("authenticated")
          // var redirectTo = resp.headers.get('X-Centinela-Redirect-To');
          if (targetUrl) {
            console.log("redirecting to ", targetUrl);
            if (!targetUrl.startsWith('http')) {
              targetUrl = "http://" + targetUrl
            }
            window.location.replace(targetUrl)
          }
        } else if (resp.status == 400) {
          resp.text().then(msg => showMessage(msg))
        } else {
          console.warn("unknown reply", resp)
        }
      },
      err => {
        console.error("err", err)
      }
    )
  }
</script>

<style>
  body {
    margin: 0px;
    padding: 0px;
    /* background: #1abc9d; */
  }

  h1 {
    color: #fff;
    text-align: center;
    font-family: Arial;
    font-weight: normal;
    margin: 2em auto 0px;
  }

  .form {
    width: 400px;
    height: 230px;
    background: #edeff1;
    margin: 200px auto;
    padding-top: 20px;
    border-radius: 10px;
    -moz-border-radius: 10px;
    -webkit-border-radius: 10px;
  }

  .text-box {
    display: block;
    width: 309px;
    height: 35px;
    margin: 15px auto;
    background: #fff;
    border: 0px;
    padding: 5px;
    font-size: 16px;
    border: 2px solid #fff;
    transition: all 0.3s ease;
    border-radius: 5px;
    -moz-border-radius: 5px;
    -webkit-border-radius: 5px;
  }

  .text-box:focus {
    border: 2px solid #1abc9d
  }

  .btn-submit {
    display: block;
    background: #1abc9d;
    width: 314px;
    padding: 12px;
    cursor: pointer;
    color: #fff;
    border: 0px;
    margin: auto;
    border-radius: 5px;
    -moz-border-radius: 5px;
    -webkit-border-radius: 5px;
    font-size: 17px;
    transition: all 0.3s ease;
  }

  .btn-submit:hover {
    background: #09cca6
  }

  #message-label {
    text-align: center;
    font-family: Arial;
    color: red;
    display: block;
    margin: 15px auto;
    text-decoration: none;
    transition: all 0.3s ease;
    font-size: 12px;
  }

  a {
    text-align: center;
    font-family: Arial;
    color: gray;
    display: block;
    margin: 15px auto;
    text-decoration: none;
    transition: all 0.3s ease;
    font-size: 12px;
  }

  a:hover {
    color: #1abc9d;
  }


  ::-webkit-input-placeholder {
    color: gray;
  }

  :-moz-placeholder {
    /* Firefox 18- */
    color: gray;
  }

  ::-moz-placeholder {
    /* Firefox 19+ */
    color: gray;
  }

  :-ms-input-placeholder {
    color: gray;
  }
</style>
`
