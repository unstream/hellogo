import React from "react";
import tinygradient from "tinygradient";
import {RGBColor} from "react-color";
import {Classes} from "reactcss";
import {ColorProp, IterationColor} from '../iteration-color/IterationColor';
import tinycolor from "tinycolor2";

export type GradientChangeHandler = (state: ColorGradientProp) => void;

export type ColorGradientProp = {
    iteration1: number,
    iteration2: number,
    color1: RGBColor,
    color2: RGBColor,
    maxIterations: number,
    onChange?: GradientChangeHandler | undefined;
};

type ColorGradientState = {
    iteration1: number,
    iteration2: number,
    color1: RGBColor,
    color2: RGBColor,
    maxIterations: number,
};

export class ColorGradient extends React.Component<ColorGradientProp, ColorGradientState>  {
    state: ColorGradientState = {
        iteration1: this.props.iteration1,
        iteration2: this.props.iteration2,
        color1: this.props.color1,
        color2: this.props.color2,
        maxIterations: this.props.maxIterations,
    };
    handleChange1 = (prop: ColorProp) => {
        this.setState({color1: tinycolor(prop.color).toRgb()})
        this.setState({iteration1: prop.iteration})
        if (this.props.onChange) {
            this.props.onChange(this.state);
        }
    };
    handleChange2 = (prop: ColorProp) => {
        this.setState({color2: tinycolor(prop.color).toRgb()})
        this.setState({iteration2: prop.iteration})
        if (this.props.onChange) {
            this.props.onChange(this.state);
        }
    };

    render() {
        //Positions: 0, color1/maxIterations, color2/maxIterations, 1
        let pos1 = this.state.iteration1 / this.state.maxIterations;
        let pos2 = this.state.iteration2 / this.state.maxIterations;
        console.log(this.state.iteration1 + "," + this.state.iteration2 + "," + this.state.maxIterations);
        if (pos2 > 1) pos2 = 1;
        const gradient = tinygradient([
            {color: this.state.color1, pos: 0},
            {color: this.state.color1, pos: pos1},
            {color: this.state.color2, pos: pos2},
            {color: this.state.color2, pos: 1}
        ]).css()

         return (
            <div>
                <div style={{background: gradient}}>
                    &nbsp;{this.state.maxIterations}
                </div>
                <IterationColor
                    iteration = {this.state.iteration1}
                    color = { tinycolor(this.state.color1).toRgb() }
                    onChange={this.handleChange1}
                />
                <IterationColor
                    iteration = {this.state.iteration2}
                    color = { tinycolor(this.state.color2).toRgb() }
                    onChange={this.handleChange2}
                />

            </div>
        )
    }
}