import * as React from 'react';
import { useNavigate } from "react-router-dom";
import { useState } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';

import { st } from './style';
import { LoginUser } from '../api/auth'

const validateEmail = (email) => {
  return String(email)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
};

export default function LayoutLogin() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const redirectToRegister = () => {
    navigate("/register");
        navigate(0)
    }


  const updateEmail = (e) => {
        setEmail(e)
  };

  const updatePassword = (e) => {
        setPassword(e)
  };

  const loginUser = () => {
    if (validateEmail(email) === null) {
            return
    }

    LoginUser(email, password).then((data) => {
      if (data.message === "invalid email or password") {
        setError("Incorrect user or Password");
      } else {
        localStorage.setItem('test-token', data.Token);
        localStorage.setItem('email', data.Email);
        navigate("/buildings");
        navigate(0);
      }
    });
  };

  return (
    <div
      style={{ width: "100vw", height: "100vh", backgroundColor: st.bgColor }}
    >
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position='static' style={{ backgroundColor: st.appBarColor }}>
          <Toolbar>
            <Typography variant='h6' component='div' sx={{ flexGrow: 1 }}>
              My Buildings
            </Typography>
            <Button color='inherit' onClick={redirectToRegister}>
              Register
            </Button>
          </Toolbar>
        </AppBar>
      </Box>
      <Grid
        container
        spacing={10}
        style={{ width: "100%", height: "calc(100% - 65px)" }}
      >
        <Grid item style={{ width: "100%", height: "20%" }}></Grid>
        <Grid item style={{ width: "100%", height: "10%" }}>
          <Typography
            variant='h6'
            component='div'
            sx={{ flexGrow: 1 }}
            style={st.center}
          >
            Login
          </Typography>
        </Grid>
        <Grid
          item
          rowSpacing={2}
          style={{ width: "100%", height: "40%", padding: "15px" }}
        >
          <Grid item style={st.center} sx={{ margin: "15px" }}>
            <TextField
              id='standard-basic'
              label='Email'
              variant='standard'
              onChange={(event) => {
                updateEmail(event.target.value);
              }}
            />
          </Grid>
          <Grid item style={st.center} sx={{ margin: "15px" }}>
            <TextField
              id='outlined-password-input'
              label='Password'
              type='password'
              autoComplete='current-password'
              variant='standard'
              onChange={(event) => {
                updatePassword(event.target.value);
              }}
            />
          </Grid>
          <Grid item style={st.center} sx={{ margin: "15px" }}>
            <Button
              variant='contained'
              onClick={loginUser}
              data-testid='login-button'
            >
              Login
            </Button>
          </Grid>
          <Typography
            variant='h6'
            component='div'
            sx={{ flexGrow: 1, color: "red" }}
            style={st.center}
            data-testid='error-message'
          >
            {error}
          </Typography>
        </Grid>
        <Grid item style={{ width: "100%", height: "30%" }}></Grid>
      </Grid>
    </div>
  );
}
