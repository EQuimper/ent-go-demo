import { Box, Button, Flex, Link } from "@chakra-ui/react";
import NextLink from "next/link";

interface Props {
  children?: React.ReactNode;
}

function Nav({ children }: Props) {
  return (
    <Box>
      <Flex
        h="60px"
        align="center"
        justify="space-between"
        px={4}
        bg="gray.100"
      >
        <Box>
          <NextLink href="/" passHref>
            <Link fontSize="xl" fontWeight="bold">
              EntDemo
            </Link>
          </NextLink>
        </Box>

        <Flex>
          <Box mr={2}>
            <NextLink href="/register" passHref>
              <Link as={Button}>Register</Link>
            </NextLink>
          </Box>
          <NextLink href="/login" passHref>
            <Link as={Button}>Login</Link>
          </NextLink>
        </Flex>
      </Flex>
      {children}
    </Box>
  );
}

export default Nav;
