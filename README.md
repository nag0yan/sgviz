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
- [ ] Multiple security groups
- [ ] Any inbound rule

## Not supported
- Outbound rules
- Image export

## License
MIT
