# ism
## 项目简介

ism保险项目后端api

主要逻辑是定时从合约中读取数据并提供该前端查询

## 启动
在配置文件中配置好相关参数，具体可参考模版
配置文件跟可执行文件放在同一目录下

```shell script
git clone https://github.com/tpkeeper/ism.git
cd ism/cmd/ismd
go build
./ismd
```

## 用 swagger 生成接口文档

在包含main方法的文件目录下执行，会自动生成接口文档

```shell script
swag init --parseDependency
```

## 生成 abi bin 文件

```shell script
abigen --abi ./RewardPools_abi.json --pkg ism_reward_pool --type IsmRewardPool --out ./ism_reward_pool.go
```

## 运维相关事项

### 添加协议前需要手动配置协议保险品地址、logo 等信息

```sql
update protocol_token_infos set logo_url = "/static/icon/ETH.png" where symbol="ETH";
update protocol_token_infos set official_url="huobitoken.com" where symbol="ht";
```

### 当 protocol token 的价格从 coinw 获取有问题时，可以调整 symbol

```sql
update protocol_token_infos set symbol = "ETH" where id=xxx;
```

### 当 product token 价格从 mdex 获取有问题时，手动设置价格

```sql
update product_token_infos set price=0.1102 where id = xxx;
```

### 当矿池lp price 获取不到时，手动设置价格

```sql
update contract_mine_pool_infos set lp_price=1.2 where id = xxx;
```
### 配置矿池 lptoken 交易链接地址
需要在 ContractMinePoolInfo 表中根据lp的名称手动配置 url 字段

ISM：无
ISM/USDT：https://ht.mdex.com/#/add/0x348ccc5a616abae8a639457fc469917b03d938c3/0xa71edc38d189767582c38a3145b5873052c3e47a
HT-CLAIM：https://ht.mdex.com/#/add/0xEC0F04f2e3066A4465770C82eC98b6de3d081Fcf/0xa71edc38d189767582c38a3145b5873052c3e47a
HT-UNCLAIM：https://ht.mdex.com/#/add/0x7CD647010ed1519102EEE52A252AAF505bd29FB1/0xa71edc38d189767582c38a3145b5873052c3e47a
ETH-CLAIM：https://ht.mdex.com/#/add/0x82B4f3b2b960746db711d69a1265299e295eBeEA/0xa71edc38d189767582c38a3145b5873052c3e47a
ETH-UNCLAIM：https://ht.mdex.com/#/add/0xA6Cf3C0bCc871d227aE31986024Bd8D53C8981a6/0xa71edc38d189767582c38a3145b5873052c3e47a