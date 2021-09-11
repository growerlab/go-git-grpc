package main

// 测试 git-upload-pack git-receive-pack 操作

func main() {
	initServer()

	if err := testPush(); err != nil {
		panic(err)
	}
}

func testPush() error {

	return nil
}

func testPull() error {
	return nil
}
