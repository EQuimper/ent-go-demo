import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { Container, Text } from "@chakra-ui/react";

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>EntDemo</title>
        <meta name="description" content="Ent demo" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <Container>
          <Text fontSize="6xl">EntDemo</Text>
        </Container>
      </main>
    </div>
  );
};

export default Home;
