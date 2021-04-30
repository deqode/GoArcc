import Head from 'next/head';
import { useRouter } from 'next/router';
import { useContext, useEffect } from 'react';
import Loading from '../components/Loading';
import { UserContext } from '../Contexts/UserContext';
import { getStorage } from '../utils/localStorage';

export default function Success() {
  const { user } = useContext(UserContext)
  const router = useRouter();
  useEffect(() => {
    if (user.state == -1)
      router.push("/")
  }, [user])
  if(user.state!=1)return (<div><Loading/></div>)
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
