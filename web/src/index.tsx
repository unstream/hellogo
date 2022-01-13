import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import theme from "./theme";
import {Container, CssBaseline, Grid, Paper, ThemeProvider, Toolbar} from "@mui/material";
import Box from '@mui/material/Box';

ReactDOM.render(
    <ThemeProvider theme={theme}>


        <Box sx={{display: 'flex'}}>
            <CssBaseline/>
            {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
            <App/>
        </Box>
    </ThemeProvider>,
    document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
