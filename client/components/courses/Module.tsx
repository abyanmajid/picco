"use client";

import React from "react";
import { Card, CardBody } from "@nextui-org/react";
import { Link } from "@nextui-org/react";
import Container from "../common/Container";
import { Divider } from "@nextui-org/react"
import { Button } from "@nextui-org/react";
import { HeartIcon } from "../ui/icons";
import { Avatar } from "@nextui-org/react";
import { Accordion, AccordionItem } from "@nextui-org/react";
import TaskCard from "./TaskCard";

import Title from "../common/Title";
import Subtitle from "../common/Subtitle";

type courseDetails = {
    title: string,
    description: string,
    creator: string,
    creatorProfileUrl: string,
    courseUrl: string,
    likes: number,
}

type module = {
    name: string,
    type: "lecture" | "exercise" | "quiz"
    xp: number,
}

type Props = {
    courseDetails: courseDetails,
}

export default function Module({ courseDetails }: Props) {
    const [userLiked, setUserLiked] = React.useState(false);
    const [likes, setLikes] = React.useState(147);
    return <>
        <Card className="bg-black border border-neutral-700 w-full text-left p-8">
            <CardBody>
                <Container className="grid grid-cols-2">
                    <Title size="sm">Variables and Data types</Title>
                    <Container className="justify-end text-right">
                        <Button variant="shadow">Return to Course</Button>
                    </Container>
                </Container>
                <div className="text-neutral-300 mt-4">Introduction to Programming</div>
                <Divider className="my-4" />
                <TaskCard taskName="Classes and Objects" taskType="lecture" isCompleted={true} />
                <TaskCard taskName="Leap Year" taskType="exercise" isCompleted={true} />
                <TaskCard taskName="Exam" taskType="quiz" isCompleted={false} />
            </CardBody>
        </Card >
    </>
}