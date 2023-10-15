import java.text.NumberFormat;

public class Main {
    public static void main(String[] args) {
        // USER DEFINED
        long initial_Investment = 40000;
        float salary = 5000f;
        float monthlySalaryShare = 0.15f; // How much the user invests, as a % of the salary
        float yield_investment = 0.10f; // Yield Rate during investment period
        float yield_dividend = 0.03f; // Yield Rate after conquering Financial Freedom (recommend between 0.02 ~0.05)
        float investment_years = 20f;
        String operation = "Yield"; // possible values: "Salary Share", "Yield", "Time", "Net Worth"
        

        final byte MONTHS_IN_YEARS = 12;
        final double NEEDED_NETWORTH = salary*MONTHS_IN_YEARS/yield_dividend; // The amount of money the user wishes to have at the end



        // Variables related to the switch/case operations
        double tentative_result = 0;
        double tentative_dividends = 0;
        final byte ITER_MAX = 25;

        switch (operation){
            case "Net Worth": // Calculating final Net Worth based on User Input
                CompoundProjector netWorthObj = new CompoundProjector(initial_Investment, salary, monthlySalaryShare, yield_investment, yield_dividend, investment_years);

                double result_end_Investment = netWorthObj.projectNetWorth();

                double monthlyDividends = netWorthObj.projectDividends();

                System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(result_end_Investment));
                System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(monthlyDividends));
                break;

            case "Yield": // Estimating Yield
                float yieldFloor = 0.0f;
                float yieldCeil = 2f;
                float yieldTentative = 0.05f;
                

                yieldCase: {
                    for(int i=0; i<=ITER_MAX; i++){
                        CompoundProjector yieldOBJ = new CompoundProjector(initial_Investment, salary, monthlySalaryShare, yieldTentative, yield_dividend, investment_years);
                        tentative_result = yieldOBJ.projectNetWorth();
                        

                        if (NEEDED_NETWORTH * 0.99999 < tentative_result && tentative_result < NEEDED_NETWORTH * 1.00001){ // Almost Equals, with 0.002% Threshold
                            tentative_dividends = yieldOBJ.projectDividends();
                            System.out.println("This calculation took " + i + " iterations.");
                            System.out.println("Estimated Yield: " + Math.round(yieldTentative * 10000)/100f + "%");
                            System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                            System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                            break yieldCase;
                        }
                        // Didn't find the correct value, 'binary search' the new yield_investment
                        if (tentative_result < NEEDED_NETWORTH){
                            yieldFloor = yieldTentative;
                            yieldTentative = (yieldFloor + yieldCeil) / 2;
                        }else{
                            yieldCeil = yieldTentative;
                            yieldTentative = (yieldFloor + yieldCeil) / 2;
                        }
                    }

                    // Didn't find a satisfactory answer but we still need to 
                    System.out.println("This calculation would take more than " + ITER_MAX + " iterations.");
                    System.out.println("Estimated Yield: " + Math.round(yieldTentative * 100f) + "%");
                    System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                    System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                }

                break;
            case "Time":
                float timeFloor = 0.0f;
                float timeCeil = 250f;
                float timeTentative = 10f;

                timeCase: {
                    for(int i=0; i<=ITER_MAX; i++){
                        CompoundProjector timeOBJ = new CompoundProjector(initial_Investment, salary, monthlySalaryShare, yield_investment, yield_dividend, timeTentative);
                        tentative_result = timeOBJ.projectNetWorth();
                        

                        if (NEEDED_NETWORTH * 0.99999 < tentative_result && tentative_result < NEEDED_NETWORTH * 1.00001){ // Almost Equals, with 0.002% Threshold
                            tentative_dividends = timeOBJ.projectDividends();
                            System.out.println("This calculation took " + i + " iterations.");
                            System.out.println("Estimated Years: " + Math.round(timeTentative * 100)/100f);
                            System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                            System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                            break timeCase;
                        }
                        // Didn't find the correct value, 'binary search' the new timeTentative
                        if (tentative_result < NEEDED_NETWORTH){
                            timeFloor = timeTentative;
                            timeTentative = (timeFloor + timeCeil) / 2;
                        }else{
                            timeCeil = timeTentative;
                            timeTentative = (timeFloor + timeCeil) / 2;
                        }
                    }

                    // Didn't find a satisfactory answer but we still need to 
                    System.out.println("This calculation would take more than " + ITER_MAX + " iterations.");
                    System.out.println("Estimated Yield: " + Math.round(timeTentative * 100f) + "%");
                    System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                    System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                }
                break;
            case "Salary Share":
                break;
        }
    }
}
