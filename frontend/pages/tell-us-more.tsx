import React, {useContext, useEffect, useState} from "react";
import {
    Button,
    Drawer,
    Fab,
    FormControl,
    Grid,
    InputLabel,
    MenuItem,
    Paper,
    Select,
    TextField,
    Typography,
} from "@material-ui/core";
import {createStyles, makeStyles, Theme} from "@material-ui/core/styles";
import {useLazyQuery, useMutation} from "@apollo/client";
import Head from "next/head";
import {useRouter} from "next/router";
import Loading from "../components/Loading";
import {UserContext} from "../Contexts/UserContext";
import {AppContext} from "../Contexts/AppContext";
import VisibilityIcon from '@material-ui/icons/Visibility';
import {CLONE_REPOSITORY, GET_BRANCHES, GET_REPOSITORIES, VCS_CONNECTIONS,} from "../GraphQL/Query";
import VerticalLinearStepper from "../components/steps";
import {LogViewer} from "react-log-output";
import PetsIcon from '@material-ui/icons/Pets';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            width: "100%",
            "& .MuiTextField-root": {
                margin: theme.spacing(1),
                width: "25ch",
            },
        },
        formControl: {
            margin: theme.spacing(2),
            minWidth: 225,
        },
        selectEmpty: {
            marginTop: theme.spacing(2),
        },
    })
);

export default function TellUsMore() {
    const classes = useStyles();
    const {user} = useContext(UserContext);
    const {app} = useContext(AppContext);
    const [ownerName, setownerName] = useState("");
    const [repos, setrepos] = useState([]);
    const [currenRepoName, setcurrenRepoName] = useState("");
    const [branches, setbranches] = useState([]);
    const router = useRouter();
    const [cloneUrl, setCloneUrl] = useState("");
    const [currentBranchName, setcurrentBranchName] = useState("");
    const [runID, setrunID] = useState("");
    const [workFlowId, setworkFlowId] = useState("");
    const [showStepper, setshowStepper] = useState(false);
    const [cloneLogs, setCloneLogs] = useState([]);
    const [cloneLog, setCloneLog] = useState("");
    const [showLogs, setshowLogs] = useState(false)


    const [vcsrefetch, vcsData] = useLazyQuery(VCS_CONNECTIONS);
    const [reposrefetch, reposData] = useLazyQuery(GET_REPOSITORIES);
    const [branchesRefetch, branchesData] = useLazyQuery(GET_BRANCHES);
    const [clonereporefetch, cloneData] = useMutation(CLONE_REPOSITORY);

    useEffect(() => {
        app.centrifuge.connect();
        app.centrifuge.subscribe("clone-logs", (message) => {
            setCloneLog(message.data.log);
            console.log("setCloneLog")
        });
    }, []);
    useEffect(() => {
        setCloneLogs([...cloneLogs, cloneLog])
        console.log("setCloneLogs")

    }, [cloneLog]);

    useEffect(() => {
        if (user.idToken != "" && user.accounts.length > 0)
            vcsrefetch({
                variables: {
                    accountid: user.accounts[0].id,
                },
            });
    }, [user]);

    useEffect(() => {
        if (
            !vcsData.loading &&
            vcsData.error == undefined &&
            vcsData.data != undefined
        )
            setownerName(vcsData.data.VCSConnections.vcs_connection[0].user_name);
    }, [vcsData]);

    useEffect(() => {
        if (ownerName != "")
            reposrefetch({
                variables: {
                    userid: user.userId,
                    accountid: user.accounts[0].id,
                    provider: user.provider,
                },
            });
    }, [ownerName]);

    useEffect(() => {
        if (
            !reposData.loading &&
            reposData.error == undefined &&
            reposData.data != undefined
        ) {
            setrepos(reposData.data.repositories.repositories);
            console.log(reposData.data.repositories.repositories, "wow");
        }
    }, [reposData]);

    useEffect(() => {
        if (currenRepoName != "")
            branchesRefetch({
                variables: {
                    ownerName: ownerName,
                    repoName: currenRepoName,
                    accountid: user.accounts[0].id,
                    provider: user.provider,
                },
            });
    }, [currenRepoName]);

    useEffect(() => {
        if (
            !reposData.loading &&
            reposData.error == undefined &&
            reposData.data != undefined
        ) {
            setrepos(reposData.data.repositories.repositories);
            console.log(reposData.data.repositories.repositories, "wow");
        }
    }, [reposData]);

    useEffect(() => {
        if (
            !branchesData.loading &&
            branchesData.error == undefined &&
            branchesData.data != undefined
        ) {
            setbranches(branchesData.data.repository.branches);
            setCloneUrl(branchesData.data.repository.clone_url);
        }
    }, [branchesData]);

    useEffect(() => {
        if (currenRepoName != "")
            branchesRefetch({
                variables: {
                    ownerName: ownerName,
                    repoName: currenRepoName,
                    accountid: user.accounts[0].id,
                    provider: user.provider,
                },
            });
    }, [currenRepoName]);

    useEffect(() => {
        if (
            !cloneData.loading &&
            cloneData.error == undefined &&
            cloneData.data != undefined
        ) {
            setrunID(cloneData.data.cloneRepository.run_id || "");
            setworkFlowId(cloneData.data.cloneRepository.workflow_id || "");
        }
    }, [cloneData]);

    const selectRepo = (e) => {
        if (e.target.value == "choose") {
            setbranches([]);
            setcurrenRepoName("");
        } else {
            setcurrenRepoName(e.target.value);
        }
    };

    const selectBranch = (e) => {
        if (e.target.value == "choose") {
            setcurrentBranchName("");
        } else {
            setcurrentBranchName(e.target.value);
        }
    };

    const cloneRepo = () => {
        if (
            user.idToken != "" &&
            user.accounts.length > 0 &&
            cloneUrl != "" &&
            user.provider != "" &&
            currenRepoName != "" &&
            currentBranchName != ""
        ) {
            clonereporefetch({
                variables: {
                    cloneURL: cloneUrl,
                    branchName: currentBranchName,
                    accountId: user.accounts[0].id,
                    provider: user.provider,
                    userName: ownerName,
                },
            });
            setshowStepper(true);
            setshowLogs(true)
        }
    };

    useEffect(() => {
        if (user.state == -1) router.push("/");
    }, [user]);
    if (user.state != 1)
        return (
            <div>
                <Loading/>
            </div>
        );
    return (
        <div>
            <Head>
                <title>Tell Us More!</title>
                <link rel="icon" href="/favicon.ico"/>
            </Head>
            <Paper elevation={0} style={{padding: "100px"}}/>
            <Grid justify="center" alignItems="center" spacing={1} container>
                <Grid item md={4}>
                    <Typography variant="h1" component="h1" style={{fontSize: "50px"}}>
                        Tell Alfed more about your project.
                    </Typography>
                </Grid>
                <Grid item>
                    <form className={classes.root} noValidate autoComplete="off">
                        <Grid container spacing={7} alignItems="center">
                            <Grid item>
                                <Typography variant="h6" align="justify" color="textSecondary">
                                    Your App's Name
                                </Typography>
                            </Grid>
                            <Grid item>
                                <TextField
                                    id="outlined-basic"
                                    label="App Name"
                                    variant="outlined"
                                    placeholder="Demo App"
                                />
                            </Grid>
                        </Grid>

                        <Grid container spacing={8} alignItems="center">
                            <Grid item>
                                <Typography variant="h6" align="justify" color="textSecondary">
                                    Your Repo's Name
                                </Typography>
                            </Grid>
                            <Grid>
                                <FormControl variant="outlined" className={classes.formControl}>
                                    <InputLabel id="demo-simple-select-outlined-label">
                                        Repo Name
                                    </InputLabel>
                                    <Select
                                        labelId="demo-simple-select-outlined-label"
                                        id="demo-simple-select-outlined"
                                        onChange={selectRepo}
                                        label="Repo Name"
                                        disabled={!(repos.length > 0)}
                                        autoWidth
                                    >
                                        <MenuItem value="">
                                            <em>Choose</em>
                                        </MenuItem>
                                        {repos.map((r) => (
                                            <MenuItem value={r.name}>{r.name}</MenuItem>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>
                        </Grid>
                        <Grid container spacing={10} alignItems="center">
                            <Grid item>
                                <Typography variant="h6" align="justify" color="textSecondary">
                                    Your Branch
                                </Typography>
                            </Grid>
                            <Grid item>
                                <FormControl variant="outlined" className={classes.formControl}>
                                    <InputLabel id="demo-simple-select-outlined-label">
                                        Branch Name
                                    </InputLabel>
                                    <Select
                                        labelId="demo-simple-select-outlined-label"
                                        id="demo-simple-select-outlined"
                                        onChange={selectBranch}
                                        label="Repo Name"
                                        disabled={!(branches.length > 0)}
                                    >
                                        <MenuItem value="">
                                            <em>Choose</em>
                                        </MenuItem>
                                        {/* {branches.map(r => <option value={r}>{r}</option>)} */}
                                        {branches.map((r) => (
                                            <MenuItem value={r}>{r}</MenuItem>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>
                            <Grid justify="center" alignItems="center" spacing={1} container>
                                <Button
                                    variant="contained"
                                    color="primary"
                                    // className={classes.button}
                                    onClick={cloneRepo}
                                    startIcon={<PetsIcon/>}
                                >
                                    Go Fetch
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Grid>
            </Grid>
            <Paper/>
            {showStepper ? (
                <div style={{width: ""}}>
                    <VerticalLinearStepper runID={runID} workflowID={workFlowId}/>
                </div>
            ) : (
                ""
            )}
            <br/>
            <Drawer anchor={"top"} open={showLogs} onClose={() => {
                setshowLogs(!showLogs)
            }}>
                <Paper style={{maxHeight: 300, overflow: 'auto'}}>
                    <LogViewer text={cloneLogs.join("\n")}/>

                </Paper>
            </Drawer>
            <Fab color="secondary" style={{
                margin: 0,
                top: 'auto',
                right: 50,
                bottom: 100,
                left: 'auto',
                position: 'fixed',
            }} onClick={() => {
                setshowLogs(!showLogs)
            }} aria-label="edit">
                <VisibilityIcon/>
            </Fab>

        </div>
    );
}
