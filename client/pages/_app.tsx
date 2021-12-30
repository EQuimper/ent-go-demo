import "../styles/globals.css";
import type { AppProps } from "next/app";
import { ChakraProvider } from "@chakra-ui/react";
import Nav from "../components/common/Nav";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <Nav>
        <Component {...pageProps} />
      </Nav>
    </ChakraProvider>
  );
}

export default MyApp;
