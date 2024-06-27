import React, { ReactNode } from "react";
import { subtitle } from "@/components/primitives";

type Props = {
    className?: string,
    children: ReactNode,
}

export default function Subtitle({ className, children }: Props) {
    return <h1 className={subtitle({ class: className })}>{children}</h1>
}