import React from "react";
import {ColorChangeHandler, ColorResult, RGBColor, SketchPicker} from "react-color";
import "./IterationColor.css";
import {Divider, IconButton, InputBase, Paper} from "@mui/material";
import {SquareSharp} from "@mui/icons-material";

export type ColorProp = {
    iteration: number,
    color: RGBColor;
    onChange?: (colorProp: ColorProp) => void;
    onChange2?: ColorChangeHandler | undefined;
};

export type ColorState = {
    iteration: number,
    color: RGBColor;
    displayColorPicker: boolean;
};

export class IterationColor extends React.Component<ColorProp, ColorState> {

    state: ColorState = {
        iteration: this.props.iteration,
        color: this.props.color,
        displayColorPicker: false,
    };


    handleClick = () => {
        this.setState({displayColorPicker: !this.state.displayColorPicker})
    };

    handleClose = () => {
        this.setState({displayColorPicker: false})
    };

    handleCompleteChange = (color: ColorResult, event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({
            color: color.rgb,
        })
        if (this.props.onChange) {
            this.props.onChange(this.state);
        }
    };

    render() {
        return (
            <div>
                <Paper variant="outlined" component="form"
                       sx={{ display: "flex", zIndex: 0, my: 1 }} >
                    <InputBase
                        sx={{ ml: 1, flex: 1 }}
                        placeholder="Iteration"
                        inputProps={{ "aria-label": "color" }}
                        type="number"
                        name={"iteration"}
                        value={this.state.iteration}

                        onChange={ (event) => {
                            this.setState({iteration: +event.target.value});
                            if (this.props.onChange) {
                                this.props.onChange({color: this.state.color, iteration: +event.target.value});
                            }
                        }}
                    />
                    <Divider sx={{ height: 24, m: 0.5 }} orientation="vertical" />
                    <IconButton sx={{
                        p: "10px",
                        color:
                            "rgb(" +
                            this.state.color.r +
                            "," +
                            this.state.color.g +
                            "," +
                            this.state.color.b +
                            ")"
                    }} aria-label="square" onClick={this.handleClick}>
                        <SquareSharp/>
                    </IconButton>
                </Paper>
                {this.state.displayColorPicker ? <div className={'popover'}>
                    <div className={'clover'} onClick={this.handleClose}/>
                    <SketchPicker color={this.state.color} onChangeComplete={this.handleCompleteChange}/>
                </div> : null}
            </div>
        )
    }
}