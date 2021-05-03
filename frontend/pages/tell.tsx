import { useLazyQuery, useQuery } from '@apollo/client';
import Head from 'next/head';
import { useRouter } from 'next/router';
import { useEffect, useState, useContext } from 'react';
import Loading from '../components/Loading';
import { UserContext } from '../Contexts/UserContext';
import { GET_BRANCHES, GET_REPOSITORIES, VCS_CONNECTIONS } from '../GraphQL/Query';

export default function Tell() {

    const { user } = useContext(UserContext)
    const [ownerName, setownerName] = useState("")
    const [repos, setrepos] = useState([])
    const [currenRepoName, setcurrenRepoName] = useState("")
    const [branches, setbranches] = useState([])
    const router = useRouter();

    const repoQuery = useLazyQuery(GET_REPOSITORIES);
    console.log(repoQuery[1].data,"data")
    useEffect(() => {
        if (user.idToken != "" && user.provider != "" && user.accounts)
            repoQuery[0]({
                variables: {
                    userid: user.userId,
                    accountid: user.accounts[0].id,
                    provider: user.provider
                }
            })
    }, [user])

    useEffect(() => {
        if (!repoQuery[1].loading && repoQuery[1].error == undefined) {
            setrepos(repoQuery[1].data.repositories.repositories)

        }
    }, [repoQuery])


    const selectRepo = (e) => {
        if (e.target.value == "choose") {
            setbranches([])
        } else
            setcurrenRepoName(e.target.value)

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