import { red } from '@mui/material/colors';
import { createTheme } from '@mui/material/styles';

// A custom theme for this app
const theme = createTheme({
    palette: {
        primary: {
            main: '#f67e06',
        },
        secondary: {
            main: '#218efc',
        },
        error: {
            main: red.A400,
        },
    },
});

export default theme;
