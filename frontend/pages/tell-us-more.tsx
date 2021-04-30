import { useQuery } from '@apollo/client';
import Head from 'next/head';
import { useRouter } from 'next/router';
import { useEffect, useState, useContext } from 'react';
import Loading from '../components/Loading';
import { UserContext } from '../Contexts/UserContext';
import { GET_BRANCHES, GET_REPOSITORIES } from '../GraphQL/Query';
import { SERVER } from '../utils/constants';

export default function TellUsMore() {


  const [accoutID, setaccoutID] = useState("")
  const { user } = useContext(UserContext)
  const [ownerName, setownerName] = useState("rjoshi-deq")
  const [repos, setrepos] = useState([])
  const [currenRepoName, setcurrenRepoName] = useState("")
  const [branches, setbranches] = useState([])

  const repoQuery = useQuery(GET_REPOSITORIES, {
    variables: {
      userid: user.userId,
      accountid: accoutID,
      provider: user.provider
    }
  });
  const reposLoading = repoQuery.loading
  const refetchRepos = repoQuery.refetch
  const repoData = repoQuery.data

  const branchQuery = useQuery(GET_BRANCHES, {
    variables: {
      ownerName: ownerName,
      repoName: currenRepoName,
      accountid: accoutID,
      provider: user.provider
    }
  })

  const branchLoading = branchQuery.loading
  const refetchBranches = branchQuery.refetch
  const branchData = branchQuery.data
  const branchError = branchQuery.error

  useEffect(() => {
    (async () => {
      if (user.userId != "") {
        let res = await fetch(`${SERVER}/account/get-user-all-account/${user.userId}`, {
          headers: new Headers({
            'Authorization': `Bearer ${user.idToken}`,
          })
        })
        let accounts = await res.json()
        setaccoutID(accounts.accounts[0].id)
        refetchRepos()
      }
    })()
  }, [user])

  useEffect(() => {
    if (!reposLoading)
      setrepos(repoData.repositories.repositories)
  }, [repoData])

  useEffect(() => {
    refetchBranches()
  }, [currenRepoName])

  console.log(branches, currenRepoName)

  useEffect(() => {
    if (!branchLoading && branchError == undefined) {
      setbranches(branchData.repository.branches)
    }
  }, [branchLoading, branchError, branchData])

  const selectRepo = (e) => {
    if (e.target.value == "choose") {
      setbranches([])
    } else
      setcurrenRepoName(e.target.value)

  }


  const router = useRouter();
  useEffect(() => {
    if (user.state == -1)
      router.push("/")
  }, [user])
  if (user.state != 1) return (<div><Loading /></div>)
  return (
    <div>
      <Head>
        <title>Tell Us More!</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
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
                  <select disabled={!(repos.length > 0)} onChange={selectRepo} className="custom-select col-md-8" id="inlineFormCustomSelectPref" defaultValue="default">
                    <option value="choose">choose</option>
                    {repos.map(r => <option value={r.name}>{r.name}</option>)}
                  </select>
                </div>
                <div className="form-group row select custom_input">
                  <label className="col-md-4" htmlFor="inlineFormCustomSelectPref">Your Branch</label>
                  <select disabled={!(branches.length > 0)} className="custom-select col-md-8" id="inlineFormCustomSelectPref" defaultValue="default">
                    {branches.map(r => <option value={r}>{r}</option>)}

                  </select>
                </div>

              </form>
              <div className="text-center mt-5 mb-5"><a href="/success" className="btn fetch-btn w-50">Go Fetch</a></div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}