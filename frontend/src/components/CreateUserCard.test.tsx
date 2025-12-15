import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';

import { CreateUserCard } from './CreateUserCard';

describe('CreateUserCard', () => {
  it('パスワードが8文字未満の場合は送信せずエラーを表示する', async () => {
    const onCreate = vi.fn().mockResolvedValue(undefined);
    const user = userEvent.setup();

    render(<CreateUserCard onCreate={onCreate} />);

    await user.type(screen.getByLabelText('ユーザー名'), 'Alice');
    await user.type(screen.getByLabelText('メールアドレス'), 'alice@example.com');
    await user.type(screen.getByLabelText('パスワード'), 'abc1234'); // 7文字

    await user.click(screen.getByRole('button', { name: '作成' }));

    expect(onCreate).not.toHaveBeenCalled();
    expect(screen.getByText('パスワードは8文字以上で入力してください')).toBeInTheDocument();
  });

  it('パスワードが半角英数字以外を含む場合は送信せずエラーを表示する', async () => {
    const onCreate = vi.fn().mockResolvedValue(undefined);
    const user = userEvent.setup();

    render(<CreateUserCard onCreate={onCreate} />);

    await user.type(screen.getByLabelText('ユーザー名'), 'Alice');
    await user.type(screen.getByLabelText('メールアドレス'), 'alice@example.com');
    await user.type(screen.getByLabelText('パスワード'), 'Password!'); // 記号あり

    await user.click(screen.getByRole('button', { name: '作成' }));

    expect(onCreate).not.toHaveBeenCalled();
    expect(screen.getByText('パスワードは半角英数字のみで入力してください')).toBeInTheDocument();
  });

  it('送信中はボタンが無効になり、成功時にフォームがクリアされる', async () => {
    const user = userEvent.setup();

    let resolve!: () => void;
    const pending = new Promise<void>((r) => {
      resolve = r;
    });
    const onCreate = vi.fn().mockReturnValue(pending);

    render(<CreateUserCard onCreate={onCreate} />);

    const name = screen.getByLabelText('ユーザー名') as HTMLInputElement;
    const email = screen.getByLabelText('メールアドレス') as HTMLInputElement;
    const password = screen.getByLabelText('パスワード') as HTMLInputElement;

    await user.type(name, 'Alice');
    await user.type(email, 'alice@example.com');
    await user.type(password, 'Password1');

    const submit = screen.getByRole('button', { name: '作成' });
    await user.click(submit);

    expect(onCreate).toHaveBeenCalledWith({
      name: 'Alice',
      email: 'alice@example.com',
      password: 'Password1',
    });

    expect(screen.getByRole('button', { name: '作成中...' })).toBeDisabled();

    resolve();

    await waitFor(() => {
      expect(screen.getByRole('button', { name: '作成' })).not.toBeDisabled();
    });

    expect(name.value).toBe('');
    expect(email.value).toBe('');
    expect(password.value).toBe('');
  });

  it('APIが失敗した場合はエラーメッセージを表示する', async () => {
    const onCreate = vi.fn().mockRejectedValue(new Error('boom'));
    const user = userEvent.setup();

    render(<CreateUserCard onCreate={onCreate} />);

    await user.type(screen.getByLabelText('ユーザー名'), 'Alice');
    await user.type(screen.getByLabelText('メールアドレス'), 'alice@example.com');
    await user.type(screen.getByLabelText('パスワード'), 'Password1');

    await user.click(screen.getByRole('button', { name: '作成' }));

    expect(await screen.findByText('boom')).toBeInTheDocument();
  });
});

