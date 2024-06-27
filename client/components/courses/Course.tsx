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
import { IoMdCheckmarkCircleOutline as CheckmarkIcon } from "react-icons/io";
import { Chip } from "@nextui-org/react";

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

export default function Course({ courseDetails }: Props) {
    const [userLiked, setUserLiked] = React.useState(false);
    const [likes, setLikes] = React.useState(147);
    return <>
        <Card className="bg-black border border-neutral-700 w-full text-left p-8">
            <CardBody>
                <Container className="grid grid-cols-2">
                    <Title size="sm">{courseDetails.title}</Title>
                    <Container className="justify-end text-right">
                        <Button
                            onClick={() => {
                                setUserLiked(!userLiked)
                                setLikes(userLiked ? likes - 1 : likes + 1)
                            }}
                            color={userLiked ? "danger" : "default"}
                            variant={userLiked ? "solid" : "ghost"}
                            aria-label="Like"
                        >
                            <span className="text-lg">{likes}</span><HeartIcon size={20} />
                        </Button>
                        <Button className="text-lg ml-2" variant="shadow" color="primary">
                            Enroll
                        </Button>
                    </Container>
                </Container>
                <Subtitle className="my-4">{courseDetails.description}</Subtitle>
                <Container className="grid grid-cols-2">
                    <Container className="flex items-center gap-4">
                        <Avatar isBordered color="primary" src="/abyan-150x150.png" />
                        <div className="flex flex-col justify-start text-left">
                            <div className="text-neutral-300">
                                Created by <Link href={courseDetails.creatorProfileUrl} className="text-primary">{courseDetails.creator}</Link>
                            </div>
                            <div className="text-neutral-300">Last updated: 06/07/24</div>
                        </div>
                    </Container>
                    <Container className="justify-end text-right">
                        <div className="text-neutral-300">26 modules</div>
                        <div className="text-neutral-300">17 students</div>
                    </Container>
                </Container>
                <Divider className="my-4" />
                <Accordion>
                    <AccordionItem key="1" subtitle="Press to expand" title="Topics (3)">
                        <ol className="list-decimal list-inside mb-4">
                            <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                            <li><span className="font-semibold">Functions:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                            <li><span className="font-semibold">Classes and objects:</span> Lorem ipsum dolor</li>
                        </ol>
                    </AccordionItem>
                    <AccordionItem key="2" subtitle="Press to expand" title="Modules (26)">
                        <Accordion variant="splitted">
                            <AccordionItem key="1" title={<span className="text-success flex items-center"><CheckmarkIcon size={20} className="mr-2" />Variables and datatypes</span>} className="text-neutral-300">
                                <ol className="list-decimal list-inside mb-4">
                                    <li className="text-success"><span className="font-semibold">Lecture:</span> Defining abstract data types in Python <Chip color="success" variant="flat">Completed</Chip></li>
                                    <li><span className="font-semibold">Quiz:</span> Debugging variables and data type errors</li>
                                    <li><span className="font-semibold">Exercise:</span> Leap year</li>
                                    <li><span className="font-semibold">Guided project:</span> Building a web server with Flask</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="2" title="Classes and objects">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                        </Accordion>
                    </AccordionItem>
                </Accordion>
            </CardBody>
        </Card >
    </>
}