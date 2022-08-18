import * as React from 'react';
import { useNavigate, useParams } from "react-router-dom";
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Moment from 'moment';

import Grid from '@mui/material/Grid';
import { st } from './style';
import { TokenIsValid } from '../api/auth'
import { GetBuildingMetrics } from '../api/buildings'
import { useState } from 'react';

import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

import TextField from '@mui/material/TextField';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { Chart } from "react-google-charts";



export default function BuildingMetrics() {
    const navigate = useNavigate();
    const [value, setValue] = React.useState(0);

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };
    const { buildingId } = useParams();
    const [data, setData] = useState([])
    const [startDate, setStartDate] = useState(new Date('2021-01-01'));
    const [endDate, setEndDate] = useState(new Date('2022-08-01'));
    const [interval, setInterval] = useState('daily')

    const [height, setHeight] = useState(0)
    const [options, setOptions] = useState({
        chart: {
            title: "Average Consumption by hour",
        },
        height: 800,
        series: {
            // Gives each series an axis name that matches the Y-axis below.
            0: { axis: "Consumption" },
            1: { axis: "Day" },
        },
        axes: {
            // Adds labels to each axis; they don't have to match the axis names.
            y: {
                Temps: { label: "Consumption (kw/h)" },
                Daylight: { label: "Day" },
            },
        },
    });

    React.useEffect(() => {
        function handleResize() {
            setHeight(window.innerHeight)
        }

        window.addEventListener("resize", handleResize)

        handleResize()
        options.height = (height - 65) * 0.79
        setOptions(options)

        return () => {
            window.removeEventListener("resize", handleResize)
        }
    }, [height])

    React.useEffect(() => {
        if (!TokenIsValid(localStorage.getItem('test-token'))) {
            navigate("/login");
            navigate(0)
        }
    }, []);

    React.useEffect(() => {
        const formattedStartDate = Moment(startDate).format('YYYY-MM-DD')
        const formattedEndDate = Moment(endDate).format('YYYY-MM-DD')
        GetBuildingMetrics(buildingId, formattedStartDate, formattedEndDate, interval).then(data => {
            const tmpData =
                [[
                    { type: "date", label: "Day" },
                    "Average energy consumption",
                ]];

            for (let I = 0; I < data.length - 1; I++) {
                const row = new Array(new Date(Date.parse(data[I].timestamp)), data[I].value)
                tmpData.push(row)
            }
            setData(tmpData)
        })


    }, [buildingId, startDate, endDate, interval]);

    const logoutUser = () => {
        localStorage.removeItem('test-token')
        navigate("/login");
        navigate(0)
    }

    const changeToDaily = () => {
        setInterval('daily')
    }

    const changeToHourly = () => {
        setInterval('hourly')
    }

    const goToBookrmarks = (_id) => {
        navigate(`/bookmarks`);
        navigate(0)
    }
    const goToBuildings = (_id) => {
        navigate(`/buildings`);
        navigate(0)
    }
    return (
        <div style={{ width: '100vw', height: '100vh', backgroundColor: st.bgColor }}>
            <Box sx={{ flexGrow: 1 }}>
                <AppBar position="static" style={{ backgroundColor: st.appBarColor }}>
                    <Toolbar>
                        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                            My Buildings
                        </Typography>
                        <Box justifyContent="flex-start" sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
                            <Button
                                onClick={goToBuildings}
                                sx={{ my: 2, color: 'white', display: 'block' }}
                            >
                                All
                            </Button>
                            <Button
                                onClick={goToBookrmarks}
                                sx={{ my: 2, color: 'white', display: 'block' }}
                            >
                                Bookmarks
                            </Button>
                        </Box>
                        {
                            localStorage.getItem('email') && <div>
                                {localStorage.getItem('email')}
                            </div>
                        }
                        <Button color="inherit" onClick={logoutUser}>Logout</Button>
                    </Toolbar>
                </AppBar>
            </Box>
            <Grid container spacing={0} style={{ width: '100%', height: 'calc(100% - 65px)' }} >
                <Grid item style={{ width: '100%', height: '15%', paddingTop: '30px' }} >
                    <LocalizationProvider dateAdapter={AdapterDateFns}>
                        <DatePicker
                            label="Start date"
                            value={startDate}
                            onChange={(newValue) => {
                                setStartDate(newValue);
                            }}
                            renderInput={(params) => <TextField {...params} />}
                        />
                    </LocalizationProvider>
                    <LocalizationProvider dateAdapter={AdapterDateFns}>
                        <DatePicker
                            label="End date"
                            value={endDate}
                            onChange={(newValue) => {
                                setEndDate(newValue);
                            }}
                            renderInput={(params) => <TextField {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item style={{ width: '100%', height: '85%', backgorundColor: 'green' }} >
                    <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                        <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
                            <Tab label="Hourly" onClick={changeToDaily} />
                            <Tab label="Daily" onClick={changeToHourly} />
                        </Tabs>
                    </Box>
                    <Chart
                        chartType="Line"
                        data={data}
                        options={options}
                    />
                </Grid>

            </Grid>
        </div>
    );
}