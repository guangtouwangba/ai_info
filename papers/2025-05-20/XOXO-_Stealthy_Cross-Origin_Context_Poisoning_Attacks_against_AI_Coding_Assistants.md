# XOXO: Stealthy Cross-Origin Context Poisoning Attacks against AI Coding Assistants

**URL**: http://arxiv.org/abs/2503.14281v2

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

当前AI编程助手广泛用于代码生成等任务，但自动收集上下文信息的过程存在安全隐患。研究者提出新型攻击方法XOXO，通过语义保留的代码修改（难以被传统分析检测）来毒化输入，诱导助手生成错误代码（成功率75.72%），还能将责任转嫁给开发者。该研究开发的GCGS黑盒攻击算法利用凯莱图系统搜索变换空间，在包括GPT-4.1和Claude 3.5等11个主流模型上验证有效。现有防御措施（如对抗性微调）对此攻击无效，凸显了AI编程工具亟需新的安全防护机制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T10:09:42Z
- **目录日期**: 2025-05-20
