# Monitoring Decomposition Attacks in LLMs with Lightweight Sequential Monitors

**URL**: http://arxiv.org/abs/2506.10949v1

## 原始摘要

Current LLM safety defenses fail under decomposition attacks, where a
malicious goal is decomposed into benign subtasks that circumvent refusals. The
challenge lies in the existing shallow safety alignment techniques: they only
detect harm in the immediate prompt and do not reason about long-range intent,
leaving them blind to malicious intent that emerges over a sequence of
seemingly benign instructions. We therefore propose adding an external monitor
that observes the conversation at a higher granularity. To facilitate our study
of monitoring decomposition attacks, we curate the largest and most diverse
dataset to date, including question-answering, text-to-image, and agentic
tasks. We verify our datasets by testing them on frontier LLMs and show an 87%
attack success rate on average on GPT-4o. This confirms that decomposition
attack is broadly effective. Additionally, we find that random tasks can be
injected into the decomposed subtasks to further obfuscate malicious intents.
To defend in real time, we propose a lightweight sequential monitoring
framework that cumulatively evaluates each subtask. We show that a carefully
prompt engineered lightweight monitor achieves a 93% defense success rate,
beating reasoning models like o3 mini as a monitor. Moreover, it remains robust
against random task injection and cuts cost by 90% and latency by 50%. Our
findings suggest that lightweight sequential monitors are highly effective in
mitigating decomposition attacks and are viable in deployment.


## AI 摘要

当前大语言模型（LLM）的安全防御在分解攻击下失效，恶意目标被拆解为无害子任务以绕过拒绝机制。现有浅层安全对齐技术仅检测即时提示的危害，无法识别长序列无害指令中的恶意意图。为此，研究提出引入外部监控器，在更高粒度上观察对话。通过构建最大规模多样化数据集（含问答、文生图、代理任务），测试显示GPT-4o平均攻击成功率高达87%。同时发现随机任务注入可进一步混淆恶意意图。提出的轻量级序列监控框架防御成功率达93%，成本降低90%，延迟减少50%，证实其高效性与部署可行性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-14T03:19:58Z
- **目录日期**: 2025-06-14
