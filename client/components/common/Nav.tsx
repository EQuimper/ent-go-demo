import { Box, Button, Flex, Link } from "@chakra-ui/react";
import NextLink from "next/link";
import { useRouter } from "next/router";
import { useQueryClient } from "react-query";
import { useAuth } from "../hooks/useAuth";

interface Props {
  children?: React.ReactNode;
}

function Nav({ children }: Props) {
  const { data, isLoading } = useAuth();
  const router = useRouter();
  const queryClient = useQueryClient();

  const logout = async () => {
    try {
      await fetch("/api/auth/logout", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
      });

      queryClient.invalidateQueries("me");
      router.push("/");
    } catch (error) {
      console.log("error", error);
    }
  };

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

        {!isLoading && (
          <>
            {!data ? (
              <Flex>
                <Box mr={2}>
                  <NextLink href="/register" passHref>
                    <Link _hover={{ textDecoration: "none" }} as={Button}>
                      Register
                    </Link>
                  </NextLink>
                </Box>
                <NextLink href="/login" passHref>
                  <Link _hover={{ textDecoration: "none" }} as={Button}>
                    Login
                  </Link>
                </NextLink>
              </Flex>
            ) : (
              <Flex>
                <Box mr={2}>
                  <NextLink href="/projects" passHref>
                    <Link _hover={{ textDecoration: "none" }} as={Button}>
                      Projects
                    </Link>
                  </NextLink>
                </Box>

                <Button onClick={logout}>Logout</Button>
              </Flex>
            )}
          </>
        )}
      </Flex>
      {children}
    </Box>
  );
}

export default Nav;
