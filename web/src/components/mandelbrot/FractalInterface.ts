import { Crop } from "react-image-crop";
import {ColorGradientProp} from "../gradient/ColorGradient";

export default interface FractalInterface {
    c0: number;
    c0i: number;
    c1: number;
    c1i: number;
    width: number;
    height: number;
    maxIterations: number;
    imageCompression: number;
    onChangeParams: (params: FractalInterface) => void;
    crop: Crop;
    gradient: ColorGradientProp;
}