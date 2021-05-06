import React, { useEffect, useRef, useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import StepContent from '@material-ui/core/StepContent';
import Box from '@material-ui/core/Box';
import Alert from '@material-ui/lab/Alert';
import Container from '@material-ui/core/Container';
import CircularProgress from '@material-ui/core/CircularProgress';
import Typography from '@material-ui/core/Typography';
import { SERVER } from '../utils/constants';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
    },
    button: {
      marginTop: theme.spacing(1),
      marginRight: theme.spacing(1),
    },
    actionsContainer: {
      marginBottom: theme.spacing(2),
    },
    resetContainer: {
      padding: theme.spacing(3),
    },
  }),
);

function getSteps() {
  return ['Repository Cloning', 'Congratulations'];
}

function getStepContent(step: number) {
  switch (step) {
    case 0:
      return `Repository Cloning Started`;
    case 1:
      return 'A Repository Cloned Successfully';
    default:
      return 'Unknown step';
  }
}

export default function VerticalLinearStepper(props: any) {
  const { workflowID, runID } = props
  const [clonningState, setclonningState] = useState(-1)
  //-1 not started yet
  //0 started
  //1 clonned
  //2 falied
  const scrollRef = useRef<HTMLDivElement>(null);

  const classes = useStyles();
  const steps = getSteps();



  useEffect(() => {
    if (workflowID != "" && runID != "") {
      setclonningState(0)
      checkClonningStatus()
    }
  }, [workflowID, runID])

  const checkClonningStatus = async (): Promise<void> => {
    console.log("checkClonningStatus")
    const res = await fetch(`${SERVER}/git-service/get-cloning-status?workflow_id=${workflowID}&runId=${runID}`)
    const data = await res.json()
    if (data.status) {
      setclonningState(1)
    } else {
      setclonningState(2)
    }
  }


  useEffect(() => {
    if (scrollRef && scrollRef.current) {
      scrollRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, []);

  return (
    <div className={classes.root} ref={scrollRef}>
      <Container fixed >
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
        >
          {clonningState == 2 ?
            <Alert severity="error">Repository Clonning Failed</Alert>
            : <Stepper activeStep={clonningState} orientation="vertical">
              {steps.map((label, index) => (
                <Step key={label}>
                  <StepLabel>{label}</StepLabel>
                  <StepContent>
                    <Typography>{getStepContent(index)}</Typography>
                    <div className={classes.actionsContainer}>
                      <div>
                        {clonningState == 0 && <CircularProgress />}
                      </div>
                    </div>
                  </StepContent>
                </Step>
              ))}
            </Stepper>}

        </Box>
      </Container>

    </div>
  );
}
