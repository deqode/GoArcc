import Head from 'next/head';

export default function Login() {
  return (
    <div>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <nav>
        <div className="container-fluid">
          <div className="logo">
            <a href="/" className="d-block"><img src="/assets/logo.png" alt="Logo!" /></a>
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
                <a href="/tell-us-more" className="btn github_btn">Login with github<img src="/assets/github_icon.png" alt="Login with github" /></a>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}
