'use client';

import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Stack,
  Text,
} from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';

const loginUsers = [{ email: 'example@gmail.com', password: 'abc12345' }];
const loginPages = ['admin', 'ds', 'de', 'biz', 'cc'];

export default function Auth() {
  const router = useRouter();
  const handleCloseModal = () => {
    router.push('/');
  };
  const schema = z.object({
    email: z
      .string()
      .min(1, { message: '必須項目です' })
      .email({ message: 'メールアドレスの形式で入力してください' }),
    password: z.string().min(1, { message: '必須項目です' }),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(schema),
    defaultValues: { email: '', password: '' },
  });
  const onValid = (data: { email: string; password: string }) => {
    const index = loginUsers.findIndex((value) => {
      return value.email == data.email && value.password == data.password;
    });
    if (index == -1) {
      router.push('/auth/failed');
    } else {
      router.push(`/workspace/${loginPages[index]}`);
    }
  };
  return (
    <>
      <Modal isOpen onClose={() => {}}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>
            <Heading size={'lg'} color={'#333333'}>
              Log In
            </Heading>
            <ModalCloseButton
              size={'lg'}
              color="#333333"
              onClick={handleCloseModal}
            />
          </ModalHeader>
          <form onSubmit={handleSubmit(onValid)}>
            <ModalBody>
              <Stack spacing={6}>
                <FormControl isInvalid={errors.email?.message !== undefined}>
                  <FormLabel color={'#333333'}>メールアドレス</FormLabel>
                  <Input {...register('email')} color={'#333333'} />
                  {errors.email?.message && (
                    <FormErrorMessage>
                      {errors.email.message.toString()}
                    </FormErrorMessage>
                  )}
                </FormControl>
                <FormControl isInvalid={errors.password?.message !== undefined}>
                  <FormLabel color={'#333333'}>パスワード</FormLabel>
                  <Input
                    type="password"
                    {...register('password')}
                    color={'#333333'}
                  />
                  {errors.password?.message && (
                    <FormErrorMessage>
                      {errors.password.message.toString()}
                    </FormErrorMessage>
                  )}
                </FormControl>
              </Stack>
            </ModalBody>

            <ModalFooter>
              <Button type="submit" colorScheme="blue">
                ログイン
              </Button>
            </ModalFooter>
          </form>
        </ModalContent>
      </Modal>
    </>
  );
}
