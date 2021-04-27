import Head from 'next/head';

export default function Success() {
  return (
    <div>
      <Head>
        <title>Tell Us More!</title>
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
              <h1 className="page-head">Alfred has cloned your app repository</h1>
            </div>
            <div className="col-md-6 my-auto login-right">
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}