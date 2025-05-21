# XOXO: Stealthy Cross-Origin Context Poisoning Attacks against AI Coding Assistants

**URL**: http://arxiv.org/abs/2503.14281v3

## 原始摘要

AI coding assistants are widely used for tasks like code generation. These
tools now require large and complex contexts, automatically sourced from
various origins$\unicode{x2014}$across files, projects, and
contributors$\unicode{x2014}$forming part of the prompt fed to underlying LLMs.
This automatic context-gathering introduces new vulnerabilities, allowing
attackers to subtly poison input to compromise the assistant's outputs,
potentially generating vulnerable code or introducing critical errors. We
propose a novel attack, Cross-Origin Context Poisoning (XOXO), that is
challenging to detect as it relies on adversarial code modifications that are
semantically equivalent. Traditional program analysis techniques struggle to
identify these perturbations since the semantics of the code remains correct,
making it appear legitimate. This allows attackers to manipulate coding
assistants into producing incorrect outputs, while shifting the blame to the
victim developer. We introduce a novel, task-agnostic, black-box attack
algorithm GCGS that systematically searches the transformation space using a
Cayley Graph, achieving a 75.72% attack success rate on average across five
tasks and eleven models, including GPT 4.1 and Claude 3.5 Sonnet v2 used by
popular AI coding assistants. Furthermore, defenses like adversarial
fine-tuning are ineffective against our attack, underscoring the need for new
security measures in LLM-powered coding tools.


## AI 摘要

AI编程助手在自动收集跨文件/项目的上下文时存在新型安全漏洞——跨源上下文投毒攻击(XOXO)。攻击者通过语义保留的代码修改(看似合法但实际有害)来污染输入，诱导AI生成错误代码(如含漏洞的代码)，成功率高达75.72%(测试含GPT-4.1等11个模型)。传统程序分析难以检测这种攻击，现有防御措施(如对抗微调)也无效。研究提出基于凯莱图的GCGS黑盒攻击算法，证明当前AI编程工具亟需新的安全防护机制。这种攻击还能将责任转嫁给受害开发者，具有隐蔽性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T02:31:49Z
- **目录日期**: 2025-05-21
