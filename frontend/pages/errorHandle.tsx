import Head from 'next/head';
import { withRouter } from 'next/router'

function ErrorHandle(props:any) {
  console.log(props)
  return (
    <div>
      <Head>
        <title>Congratulation !!</title>
        <link rel="icon" href="assets/alfred.png" />
      </Head>
      <section className="content_section">
        <div className="container">
          <div className="row">
            <div className="col-md-6">
              <h1 className="page-head">Alfred has found {props.router.query.error || "Some Error"}</h1>
            </div>
            <div className="col-md-6 my-auto login-right">
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

export default withRouter(ErrorHandle)
