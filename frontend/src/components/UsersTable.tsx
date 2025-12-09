export const UsersTable = () => {
    return (
        <div className="users-table">
            <table>
                <thead>
                    <tr>
                        <th>ユーザー名</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>ユーザー名</td>
                        <td>メールアドレス</td>
                        <td>ステータス</td>
                        <td>削除</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}