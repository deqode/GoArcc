import '../../css/bootstrap.min.css';
import '../../css/alfred.css';
import '../../css/fonts.css';
import github from '../../images/github_icon.png';
import logo from '../../images/logo.png';


function Login() {

  const signUp = () => {

    fetch("http://localhost:8082/v1/authentication/login").then(res => {
      console.log(res);
    })
    console.log("SignUp")
  }

  return (
    <div>
      <nav>
        <div className="container-fluid">
          <div className="logo">
            <a href="/" className="d-block"><img src={logo} alt="Logo!" /></a>
          </div>
        </div>
      </nav>

      <section className="content_section">
        <div className="container">
          <div className="row">
            <div className="col-md-6">
              <h1 className="page-head">Let us make your cloud work for you</h1>
              <p className="content-para">provide your repo, select your cloud, and let Alfred do the heavy lifting.</p>
            </div>
            <div className="col-md-6 login-right">
              <div className="text-center">
                <div className="sign_up_head">Sign Up</div>
                <button onClick={signUp} className="btn github_btn">Login with github<img src={github} alt="Login with github" /></button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default Login
