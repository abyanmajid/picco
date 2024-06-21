import React, { ReactNode } from "react";

interface Props {
    className?: string;
    children: ReactNode;
}

export default function Container({ className = '', children }: Props) {
    return <div className={className}>{children}</div>;
}
