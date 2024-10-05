# sgviz
A Visualizer of aws security groups

## Usage
Get security groups information in your account with [AWS CLI](https://docs.aws.amazon.com/cli/)  
```shell
aws ec2 describe-security-groups > sg.json
```
Export to mermaid graph  
```
sgviz sg.json > sg.md
```

## Supported
Any type of inbound rules
- [x] From IPv4s
- [x] From Security Groups
- [x] From Prefix Lists
- [x] From Ipv6

## Not supported
- Outbound rules
- Image export

## License
MIT
