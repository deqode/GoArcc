import { red } from '@material-ui/core/colors'
import { createMuiTheme } from '@material-ui/core/styles'

// Create a theme instance.
const theme = createMuiTheme({
  typography: {
    fontFamily: ['roboto'].join(','),
  },
  palette: {
    primary: {
      main: '#271e78',
    },
    secondary: {
      main: '#453215',
    },
    error: {
      main: red.A400,
    },
    background: {
      default: '#fff',
    },
  },
})

export default theme
