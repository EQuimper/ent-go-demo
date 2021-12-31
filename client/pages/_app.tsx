import "../styles/globals.css";
import type { AppProps } from "next/app";
import { Box, ChakraProvider } from "@chakra-ui/react";
import { QueryClient, QueryClientProvider } from "react-query";
import Nav from "../components/common/Nav";

const queryClient = new QueryClient();

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <ChakraProvider>
        <Nav>
          <Box minHeight="100vh">
            <Component {...pageProps} />
          </Box>
        </Nav>
      </ChakraProvider>
    </QueryClientProvider>
  );
}

export default MyApp;
