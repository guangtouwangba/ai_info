# ASP-based Multi-shot Reasoning via DLV2 with Incremental Grounding

**URL**: http://arxiv.org/abs/2412.17143v4

## 原始摘要

DLV2 is an AI tool for Knowledge Representation and Reasoning which supports
Answer Set Programming (ASP) - a logic-based declarative formalism,
successfully used in both academic and industrial applications. Given a logic
program modelling a computational problem, an execution of DLV2 produces the
so-called answer sets that correspond one-to-one to the solutions to the
problem at hand. The computational process of DLV2 relies on the typical Ground
&amp; Solve approach where the grounding step transforms the input program into a
new, equivalent ground program, and the subsequent solving step applies
propositional algorithms to search for the answer sets. Recently, emerging
applications in contexts such as stream reasoning and event processing created
a demand for multi-shot reasoning: here, the system is expected to be reactive
while repeatedly executed over rapidly changing data. In this work, we present
a new incremental reasoner obtained from the evolution of DLV2 towards iterated
reasoning. Rather than restarting the computation from scratch, the system
remains alive across repeated shots, and it incrementally handles the internal
grounding process. At each shot, the system reuses previous computations for
building and maintaining a large, more general ground program, from which a
smaller yet equivalent portion is determined and used for computing answer
sets. Notably, the incremental process is performed in a completely transparent
fashion for the user. We describe the system, its usage, its applicability and
performance in some practically relevant domains. Under consideration in Theory
and Practice of Logic Programming (TPLP).


## AI 摘要

DLV2是一个基于逻辑的知识表示与推理AI工具，支持答案集编程（ASP）。它采用"Ground & Solve"方法，将逻辑程序转换为等价的基础程序后求解答案集。为适应流推理等新兴应用需求，DLV2发展出增量推理功能：系统在多次执行时保持运行状态，复用先前计算构建更通用的基础程序，每次仅处理等效子集来高效生成答案集。这一增量过程对用户完全透明。研究展示了该系统在实际领域的适用性和性能表现。论文已被《逻辑编程理论与实践》(TPLP)期刊收录评审。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T02:27:54Z
- **目录日期**: 2025-04-03
