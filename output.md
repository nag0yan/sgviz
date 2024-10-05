```mermaid
flowchart LR
	sg-09fd25fe9c1abaea3["sg-09fd25fe9c1abaea3
(test-sg-1)"]
	8.8.8.8/32["8.8.8.8/32
(Dummy IP)"]
	8.8.8.8/32-->|"80-80"|sg-09fd25fe9c1abaea3
	sg-058de1ab2168c9bb6["sg-058de1ab2168c9bb6
(test-sg-3)"]
	2001:4860:4860::8888/128["2001:4860:4860::8888/128
(Dummy IP)"]
	2001:4860:4860::8888/128-->|"80-80"|sg-058de1ab2168c9bb6
	sg-0a86b346f3fc18b8a["sg-0a86b346f3fc18b8a
(test-sg-2)"]
	8.8.8.8/32["8.8.8.8/32
(Dummy IP)"]
	8.8.8.8/32-->|"80-80"|sg-0a86b346f3fc18b8a
	sg-09fd25fe9c1abaea3-->|"0-65535"|sg-0a86b346f3fc18b8a
	pl-00ee65e523c87193e["pl-00ee65e523c87193e
(Dummy Prefix List)"]
	pl-00ee65e523c87193e-->|"53-53"|sg-0a86b346f3fc18b8a
	sg-09db6769ea6dc4ade["sg-09db6769ea6dc4ade
(default)"]
	sg-09db6769ea6dc4ade-->|"0-0"|sg-09db6769ea6dc4ade
```