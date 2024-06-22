import React, { ReactNode } from "react";
import { title } from "@/components/primitives";

type Props = {
    color?: "violet" | "yellow" | "blue" | "cyan" | "green" | "pink" | "foreground" | undefined,
    size?: "sm" | "md" | "lg" | undefined,
    className?: string,
    children: ReactNode,
}

export default function Title({ color, size, className, children }: Props) {
    return <h1 className={title({ color: color, size: size, className: className })}>{children}</h1>;
}