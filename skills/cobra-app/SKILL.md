---
name: cobra-app
description: 在 daxe 项目中添加新的 Cobra CLI 命令，严格遵循 CMD/Logic 分层架构。当用户要求添加新命令、新增子命令、创建新的命令模块时触发。适用于涉及 cmd/ 和 internal/logic/ 目录的命令开发工作。
---

# Cobra 命令开发

基于分层架构的 Cobra CLI 命令开发指导，严格遵循 CMD/Logic 分离、单文件原则、配置驱动。

## 架构约束

- **cmd/ 根目录只有 root.go 一个文件**
- **CMD 层**：命令定义、参数绑定、参数验证、用户交互
- **Logic 层**：纯业务逻辑，不引用 cobra 或 CMD 包
- **两层通过 `*Config` 结构体交互**，单向依赖

## 目录结构

```
cmd/
├── root.go                     # 唯一根文件，注册主命令
├── [module]/                   # 命令模块目录
│   ├── root.go                # GetXxxCommand() 工厂函数
│   ├── [subcommand].go         # 子命令定义
│   └── validate.go             # 模块参数验证
internal/logic/
├── [module]/
│   ├── [subcommand].go         # 子命令业务逻辑（单文件原则）
│   └── common.go               # 模块通用功能（可选）
```

## 开发流程

### 1. CMD 层 — 创建子命令文件

文件：`cmd/[module]/[subcommand].go`

```go
package [module]

var subParam string

func getSubCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "sub",
        Short: "简短描述",
        Long: `详细描述。

支持的操作:
  - 功能说明

使用方式:
  daxe [module] sub ./file              # 示例

支持参数:
  -p, --param     参数说明`,
        Run: runSubCommand,
    }
    cmd.Flags().StringVarP(&subParam, "param", "p", "", "参数描述")
    return cmd
}

func runSubCommand(cmd *cobra.Command, args []string) {
    // 1. 参数验证
    if err := validateSubParams(args, subParam); err != nil {
        fmt.Printf("❌ 参数验证失败: %v\n", err)
        cmd.Help()
        return
    }

    // 2. 构建配置
    config := &sub.SubConfig{
        ThreadCount: subThreadCount,
        InputPath:   args[0],
    }

    // 3. 创建处理器并执行
    processor := sub.NewSubProcessor(config, common.AppConfigModel)
    if _, err := processor.Execute(context.Background()); err != nil {
        fmt.Printf("❌ 操作失败: %v\n", err)
    }
}
```

### 2. CMD 层 — 注册命令

文件：`cmd/[module]/root.go` — 在 `GetXxxCommand()` 中添加子命令

```go
func GetXxxCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "module",
        Short: "模块描述",
        Long:  moduleDescription,
    }
    cmd.AddCommand(getSubCommand())  // 注册子命令
    return cmd
}
```

### 3. CMD 层 — 添加验证

文件：`cmd/[module]/validate.go`

```go
func validateSubParams(args []string, param string) error {
    if len(args) == 0 {
        return fmt.Errorf("必须指定输入路径")
    }
    return nil
}
```

### 4. Logic 层 — 实现业务逻辑

文件：`internal/logic/[module]/sub.go`（单文件原则）

```go
package [module]

type SubConfig struct {
    ThreadCount int
    InputPath   string
}

type SubProcessor struct {
    config    *SubConfig
    appConfig *models.AppConfig
}

func NewSubProcessor(config *SubConfig, appConfig *models.AppConfig) *SubProcessor {
    return &SubProcessor{config: config, appConfig: appConfig}
}

func (p *SubProcessor) Execute(ctx context.Context) (*SubResult, error) {
    // 纯业务逻辑，不依赖 CMD 层
    return nil, nil
}
```

### 5. 注册主命令

文件：`cmd/root.go` — 在 init() 中注册

```go
rootCmd.AddCommand([module].GetXxxCommand())
```

## 关键模式

### 配置传递

```go
// CMD 层构建配置 → Logic 层消费
config := &md.MDConfig{ThreadCount: threads}
processor := md.NewMDFixer(config, common.AppConfigModel)
```

### 并发处理

```go
// channel + goroutine + WaitGroup
func (p *Processor) processConcurrently(tasks []Task) error {
    var wg sync.WaitGroup
    taskChan := make(chan Task, len(tasks))
    for i := 0; i < p.config.ThreadCount; i++ {
        wg.Add(1)
        go func() { defer wg.Done(); for t := range taskChan { p.process(t) } }()
    }
    for _, t := range tasks { taskChan <- t }
    close(taskChan)
    wg.Wait()
}
```

### 原子文件写入

```go
utils.WriteFileAtomically(path, []byte(content))  // []byte
md.WriteFileAtomically(path, content)                // string
```

### 文件列表加载

```go
paths, err := md.LoadFileList(listPath)  // 支持 JSON/TXT
```

## 命名规范

- 参数变量：`[模块前缀][参数名]`（如 `subInputList`）
- 验证函数：`validate[CommandName]Params`
- 工厂函数：`get[CommandName]Command`（小写，不导出）
- 主命令工厂：`GetXxxCommand`（大写，导出）
- Logic 构造器：`New[Name]Processor(config, appConfig)`
- Logic 入口：`Execute(ctx context.Context) (*Result, error)`

## 检查清单

### CMD 层
- [ ] cmd/ 根目录只有 root.go
- [ ] 命令命名简洁明了（单数形式）
- [ ] 参数命名一致性（前缀风格统一）
- [ ] 帮助信息详细有用（含使用示例）
- [ ] 验证逻辑放在 validate.go

### Logic 层
- [ ] 不引用 cobra 或 CMD 包
- [ ] 两层通过 Config 结构体交互
- [ ] 子命令逻辑在单个 .go 文件中（单文件原则）
- [ ] 使用工厂方法创建处理器
- [ ] 纯函数优先

### 架构
- [ ] 分层清晰（CMD/Logic 职责明确）
- [ ] 避免过度抽象
- [ ] 错误处理用户友好
