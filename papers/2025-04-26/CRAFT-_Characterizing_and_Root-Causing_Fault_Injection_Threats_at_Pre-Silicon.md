# CRAFT: Characterizing and Root-Causing Fault Injection Threats at Pre-Silicon

**URL**: http://arxiv.org/abs/2503.03877v2

## 原始摘要

Fault injection attacks represent a class of threats that can compromise
embedded systems across multiple layers of abstraction, such as system
software, instruction set architecture (ISA), microarchitecture, and physical
implementation. Early detection of these vulnerabilities and understanding
their root causes, along with their propagation from the physical layer to the
system software, is critical to secure the cyberinfrastructure. This work
presents a comprehensive methodology for conducting controlled fault injection
attacks at the pre-silicon level and an analysis of the underlying system for
root-causing behavior. As the driving application, we use the clock glitch
attacks in AI/ML applications for critical misclassification. Our study aims to
characterize and diagnose the impact of faults within the RISC-V instruction
set and pipeline stages, while tracing fault propagation from the circuit level
to the AI/ML application software. This analysis resulted in discovering two
new vulnerabilities through controlled clock glitch parameters. First, we
reveal a novel method for causing instruction skips, thereby preventing the
loading of critical values from memory. This can cause disruption and affect
program continuity and correctness. Second, we demonstrate an attack that
converts legal instructions into illegal ones, thereby diverting control flow
in a manner exploitable by attackers. Our work underscores the complexity of
fault injection attack exploits and emphasizes the importance of preemptive
security analysis.


## AI 摘要

这篇论文提出了一种在芯片设计前阶段进行可控故障注入攻击的综合方法，重点分析了RISC-V架构在AI/ML应用中的时钟毛刺攻击。研究发现两种新漏洞：1）通过特定时钟毛刺参数可导致指令跳过，阻止关键数据从内存加载，破坏程序连续性和正确性；2）能将合法指令转为非法指令，从而被攻击者利用控制流。研究揭示了故障从电路层到系统软件的传播路径，强调了在芯片设计阶段进行预防性安全分析的重要性，以保护网络基础设施免受这类跨抽象层攻击的威胁。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-26T20:02:05Z
- **目录日期**: 2025-04-26
