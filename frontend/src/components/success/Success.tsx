import '../../css/bootstrap.min.css'
import '../../css/alfred.css';
import '../../css/fonts.css';
import logo from '../../images/logo.png';

function Success() {
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
              <h1 className="page-head">Alfred has cloned your app repository</h1>
            </div>
            <div className="col-md-6 my-auto login-right">
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default Success
