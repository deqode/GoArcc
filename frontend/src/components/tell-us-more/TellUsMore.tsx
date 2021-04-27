import '../../css/bootstrap.min.css'
import '../../css/alfred.css';
import '../../css/fonts.css';
import logo from '../../images/logo.png';

function TellUsMore() {
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
            <div className="col-md-12 project-info_form">
              <h1 className="page-head">Tell Alfed more about your project.</h1>
              <form className="form container">
                <div className="form-group row custom_input custom_input">
                  <label htmlFor="appname" className="col-md-4">Your App's Name</label>
                  <input type="text" className="form-control col-md-8" placeholder="Demo App" />
                </div>
                <div className="form-group row select custom_input">
                  <label className="col-md-4" htmlFor="inlineFormCustomSelectPref">Your Repo's Name</label>
                  <select className="custom-select col-md-8" id="inlineFormCustomSelectPref" defaultValue="default">
                    <option value="default">email finder</option>
                    <option value="1">One</option>
                    <option value="2">Two</option>
                    <option value="3">Three</option>
                  </select>
                </div>
                <div className="form-group row select custom_input">
                  <label className="col-md-4" htmlFor="inlineFormCustomSelectPref">Your Environment</label>
                  <select className="custom-select col-md-8" id="inlineFormCustomSelectPref" defaultValue="default">
                    <option value="default">Production</option>
                    <option value="1">One</option>
                    <option value="2">Two</option>
                    <option value="3">Three</option>
                  </select>
                </div>
                <div className="form-group row custom_input">
                  <label htmlFor="branchname" className="col-md-4">Your Branch</label>
                  <input type="text" className="form-control col-md-8" placeholder="master" />
                </div>
              </form>
              <div className="text-center mt-5 mb-5"><a href="/success" className="btn fetch-btn w-50">Go Fetch</a></div>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default TellUsMore
