import { useLazyQuery, useQuery } from '@apollo/client';
import Head from 'next/head';
import { useRouter } from 'next/router';
import { useEffect, useState, useContext } from 'react';
import Loading from '../components/Loading';
import { UserContext } from '../Contexts/UserContext';
import { GET_BRANCHES, GET_REPOSITORIES, VCS_CONNECTIONS } from '../GraphQL/Query';

export default function TellUsMore() {

    const { user } = useContext(UserContext)
    const [ownerName, setownerName] = useState("")
    const [repos, setrepos] = useState([])
    const [currenRepoName, setcurrenRepoName] = useState("")
    const [branches, setbranches] = useState([])
    const router = useRouter();

    const [vcsrefetch, vcsData] = useLazyQuery(VCS_CONNECTIONS)
    const [reposrefetch, reposData] = useLazyQuery(GET_REPOSITORIES)
    const [branchesRefetch, branchesData] = useLazyQuery(GET_BRANCHES)

    useEffect(() => {
        if (user.idToken != "" && user.accounts.length > 0)
            vcsrefetch({
                variables: {
                    accountid: user.accounts[0].id
                }
            })

    }, [user])

    useEffect(() => {
        if (!vcsData.loading && vcsData.error == undefined && vcsData.data != undefined)
            setownerName(vcsData.data.VCSConnections.vcs_connection[0].user_name)

    }, [vcsData])

    useEffect(() => {
        if (ownerName != "")
            reposrefetch({
                variables: {
                    userid: user.userId,
                    accountid: user.accounts[0].id,
                    provider: user.provider
                }
            })
    }, [ownerName])


    useEffect(() => {
        if (!reposData.loading && reposData.error == undefined && reposData.data != undefined) {
            setrepos(reposData.data.repositories.repositories)
            console.log(reposData.data.repositories.repositories, "wow")
        }

    }, [reposData])

    useEffect(() => {
        if (currenRepoName != "")
            branchesRefetch({
                variables: {
                    ownerName: ownerName,
                    repoName: currenRepoName,
                    accountid: user.accounts[0].id,
                    provider: user.provider
                }
            })
    }, [currenRepoName])


    useEffect(() => {
        if (!branchesData.loading && branchesData.error == undefined && branchesData.data != undefined) {
            setbranches(branchesData.data.repository.branches)
        }
    }, [branchesData])

    const selectRepo = (e) => {
        if (e.target.value == "choose") {
            setbranches([])
            setcurrenRepoName("")

        } else {
            setcurrenRepoName(e.target.value)
        }

    }

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
                            <h5>{ownerName}</h5>
                            <form className="form container">
                                <div className="form-group row custom_input custom_input">
                                    <label htmlFor="appname" className="col-md-4">Your App's Name </label>
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
