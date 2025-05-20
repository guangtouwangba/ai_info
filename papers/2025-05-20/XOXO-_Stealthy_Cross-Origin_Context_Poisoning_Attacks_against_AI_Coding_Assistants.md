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

研究人员提出了一种新型攻击方法XOXO（跨源上下文投毒），通过语义等效的对抗性代码修改来毒化AI编程助手的输入上下文，导致其生成错误输出。该方法利用Cayley图系统搜索变换空间（GCGS算法），在11个模型（包括GPT-4.1和Claude 3.5）上平均攻击成功率达75.72%。由于修改保持代码语义正确，传统程序分析难以检测，攻击者可将责任转嫁给开发者。现有防御措施（如对抗性微调）对此攻击无效，凸显了LLM编程工具需要新的安全防护机制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T06:02:29Z
- **目录日期**: 2025-05-20
