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

AI编码助手广泛用于代码生成等任务，但其自动收集的跨文件/项目的上下文可能被恶意篡改。研究者提出新型攻击XOXO，通过语义保留的代码修改（传统分析难以检测）毒化上下文，操纵AI生成错误代码（成功率75.72%），还能嫁祸开发者。攻击算法GCGS利用凯莱图搜索变换空间，对GPT-4.1等11个模型有效。现有防御（如对抗微调）无效，凸显LLM编码工具亟需新的安全措施。该攻击具有任务无关性、黑盒特性，影响包括Claude 3.5等主流助手。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T23:02:20Z
- **目录日期**: 2025-05-20
